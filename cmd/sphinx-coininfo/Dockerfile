FROM jrei/systemd-centos:7

COPY SphinxCoininfo.viper.yaml /etc/SphinxCoininfo/SphinxCoininfo.viper.yaml
COPY sphinx-coininfo /usr/local/bin/sphinx-coininfo

CMD ["sphinx-coininfo", "run"]
