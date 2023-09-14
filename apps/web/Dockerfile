# Use an official Node.js runtime as a parent image
FROM node:18-alpine

# Set the working directory in the container
WORKDIR /app

# Copy the built Next.js project into the container
COPY dist/apps/web /app

# Install production dependencies only
RUN npm install --production

# Expose the port that your Next.js app will run on
EXPOSE 4200

# Define the command to run your Next.js app
CMD ["npm", "start"]
