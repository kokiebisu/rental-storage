export default function ImageUploader({ uploadURL, handleImageChange }) {
  return (
    <div>
      <h3>Upload Image</h3>
      <input type="file" onChange={handleImageChange} />
      <br />
      <br />
      <br />
      {uploadURL && (
        <div>
          Upload URL: <a href={uploadURL.url}>{uploadURL.url}</a>
          <br />
          <br />
        </div>
      )}
    </div>
  );
}
