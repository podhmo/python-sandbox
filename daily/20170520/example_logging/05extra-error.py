import logging
logger = logging.getLogger(__name__)


def main():
    logger.info("ok")


if __name__ == "__main__":
    # fmt = "%(asctime)s: level=%(levelname)s %(filename)s:%(lineno)s -- %(message)s"
    fmt = "%(asctime)s: me=%(me)s level=%(levelname)s %(filename)s:%(lineno)s -- %(message)s"
    logging.basicConfig(level=logging.INFO, format=fmt)
    main()
