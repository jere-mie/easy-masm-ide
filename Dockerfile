# Use the base image
FROM jeremiedevelops/winego

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Expose port 8080
EXPOSE 8080

# Run the application
CMD ["go", "run", "main.go"]
