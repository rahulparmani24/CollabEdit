# Real-Time Collaborative Editor - Frontend

This is the frontend for the Real-Time Collaborative Editor, designed to be hosted on **AWS S3** for global availability.

## Deployment Steps

1. **Build the Frontend**:
   - Ensure all files (`index.html`, `main.js`, `styles.css`) are ready.

2. **Create an S3 Bucket**:
   - Go to the AWS S3 console and create a new bucket.
   - Enable static website hosting for the bucket.

3. **Upload Files**:
   - Upload `index.html`, `main.js`, and `styles.css` to the bucket.

4. **Set Permissions**:
   - Make the bucket and its contents public so that users can access the frontend.

5. **Access the Frontend**:
   - Use the S3 static website URL (e.g., `https://your-bucket-name.s3.amazonaws.com`).

## Environment Variables

- **`BACKEND_API_URL`**: The WebSocket URL for the backend server.
  - Example: `ws://your-backend-url/ws`
