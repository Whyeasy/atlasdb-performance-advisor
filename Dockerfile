FROM alpine

COPY atlasdb-performance-advisor /usr/bin/
ENTRYPOINT ["/usr/bin/atlasdb-performance-advisor"]