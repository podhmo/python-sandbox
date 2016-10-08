import time
import random
import asyncio
import logging


logger = logging.getLogger(__name__)


class MockRequest:
    def __init__(self, domain, *args):
        self.domain = domain
        self.args = args
        self.links = []

    def add_link(self, request):
        self.links.append(request)
        return self

    def get_links(self):
        return self.links


async def fetch(request):
    d = 0.2 * random.random()
    logger.info("r start: %s args=%s, cost=%s", request.domain, request.args, d)
    await asyncio.sleep(d)
    logger.info("r end  : %s arg=%s", request.domain, request.args)
    return request.get_links()


async def tick(d):
    while True:
        logger.info("--------------------tick %s--------------------", d)
        await asyncio.sleep(d)


class Crawler:
    def __init__(self):
        self.running_count = 0
        self.cache = {}  # domain -> requester

        self.total_request = 0
        self.start_time = time.time()

    def crawl(self, request, fn, *args, **kwargs):
        self.running_count += 1
        self.total_request += 1

        domain = request.domain
        requester = self.cache.get(domain)
        if requester is None:
            self.cache[domain] = requester = self.create_requester(domain)
        fut = asyncio.ensure_future(requester(request, fn, *args, **kwargs))
        fut.add_done_callback(self._dec_count)
        return fut

    def _dec_count(self, fut):
        self.running_count -= 1

    def create_requester(self, domain):
        return Requester(domain)

    def is_running(self):
        return self.running_count > 0

    def on_finished(self):
        logger.info("end: takes %s, total request = %s", time.time() - self.start_time, self.total_request)


class Requester:
    def __init__(self, domain, limit=2):
        self.domain = domain
        self.semaphore = asyncio.Semaphore(limit)

        self.limit = limit
        self.delta = 1
        self.waittime = 1

        self.count = 0
        self.last_called_at = None

    async def __call__(self, request, fn, *args, **kwargs):
        async with self.semaphore:
            logger.info("c request: %s args=%s", self.domain, request.args)
            t = time.time()
            if self.count <= self.limit and self.last_called_at and (t - self.last_called_at) < self.delta:
                await self.wait()
            self.last_called_at = t
            self.count += 1
            result = await fn(request, *args, **kwargs)
            self.count -= 1
            return result

    async def wait(self):
        logger.info("c *wait**: %s, waittime=%s", self.domain, self.waittime)
        await asyncio.sleep(self.waittime)


async def do_loop(requests, inq, outq, crawler):
    outq.put_nowait([crawler.crawl(req, fn=fetch) for req in requests])
    while crawler.is_running():
        if not outq.empty():
            futs = await outq.get()
            if futs:
                finished, not_finished = await asyncio.wait(futs, return_when=asyncio.FIRST_COMPLETED)
                for fut in finished:
                    inq.put_nowait(fut.result())
                outq.put_nowait(not_finished)
        if not inq.empty():
            returned_requests = await inq.get()
            futs = [crawler.crawl(req, fn=fetch) for req in returned_requests]
            if futs:
                outq.put_nowait(futs)


if __name__ == "__main__":
    logging.basicConfig(format="%(asctime)s %(message)s", level=logging.DEBUG)
    loop = asyncio.get_event_loop()

    R = MockRequest
    requests = [
        R("x", 1).add_link(R("xy", 1)).add_link(R("xy", 2)).add_link(R("xy", 3)),
        R("y", 1),
        R("y", 2),
        R("y", 11).add_link(R("yz", 1)).add_link(R("yz", 2)).add_link(R("yz", 3)),
        R("y", 12),
        R("x", 2),
        R("y", 3),
        R("z", 1),
        R("x", 3),
    ]
    inq = asyncio.Queue()
    outq = asyncio.Queue()
    crawler = Crawler()

    # # ticker
    # asyncio.ensure_future(tick(1))

    loop.run_until_complete(do_loop(requests, inq, outq, crawler))
    crawler.on_finished()
    loop.close()
