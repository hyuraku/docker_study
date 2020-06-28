# Use node 4.4.5 LTS
FROM node:12

# Copy source code
COPY . /app

# Change working directory
WORKDIR /app

# Install dependencies
RUN npm install && npm upgrage

# Expose API port to the outside
EXPOSE 80

# Launch application
CMD ["npm","start"]