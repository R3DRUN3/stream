# Use Alpine Linux as base image
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy custom module file into the container
COPY test_script.sh .

# Make the script executable
RUN chmod +x test_script.sh

# Command to run the script
CMD ["./test_script.sh"]
