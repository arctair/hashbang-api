FROM debian
COPY bin/hashbang /bin/hashbang
ENTRYPOINT ["/bin/hashbang"]
