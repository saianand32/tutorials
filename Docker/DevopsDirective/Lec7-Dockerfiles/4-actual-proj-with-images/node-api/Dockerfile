# To simplify your Dockerfile and just use the official Node.js image, 
# you don't need to manually install nodejs and npm using apt because the official Node.js 
# image (like node:19.6-bullseye-slim) already comes with Node.js and npm pre-installed.
# So, you can combine everything into a single FROM directive and avoid the unnecessary apt update/install step.
# Pin specific version for stability
# Use slim for reduced image size 
FROM node:19.6-bullseye-slim 

# Set NODE_ENV
ENV NODE_ENV production

WORKDIR /usr/src/app 

COPY ./api-node/package*.json .

# Install only production dependencies
# Use cache mount to speed up install of existing dependencies
RUN --mount=type=cache,target=/usr/src/app/.npm \
  npm set cache /usr/src/app/.npm && \
  npm ci --only=production

# Use non-root user
USER node

# Use --chown on COPY commands to set file permissions for the user
COPY --chown=node:node ./api-node/src .

CMD ["node", "index.js"]