FROM pactfoundation/pact-stub-server:latest

COPY pact /pact

RUN echo '#!/bin/sh' > /entrypoint.sh
RUN echo 'set -e' >> /entrypoint.sh
RUN echo 'export PATH=$PATH:/app' >> /entrypoint.sh
RUN echo 'pact-stub-server' \
         '-p 8080' \
         '-d /pact/service-user' \
         '-d /pact/service-space' \
         >> /entrypoint.sh 
RUN chmod +x /entrypoint.sh

ENTRYPOINT [ "/entrypoint.sh" ]