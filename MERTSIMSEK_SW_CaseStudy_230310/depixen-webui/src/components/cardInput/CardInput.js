import React from "react";
import "./CardInput.css";

function CardInput({
  handleClick,
  title,
  setTitle,
  desc,
  setDesc,
  img,
  onImageSelection,
  buttonActive,
  handleSubmit,
}) {
  return (
    <div>
      <h1 className="title-above-card">New Title</h1>
      <form onSubmit={handleSubmit} className="card-for-input">
        <input
          type="text"
          wrap="soft"
          value={title}
          className="title-in-card"
          onClick={handleClick}
          onChange={(e) => {
            setTitle(e.target.value);
          }}
        />
        <textarea
          value={desc}
          className="description-in-card"
          onClick={handleClick}
          onChange={(e) => {
            setDesc(e.target.value);
          }}
        />
        {img == null ? (
          <>
            <label htmlFor="img-upload" className="image-area">
              <p className="image-upload-sign">+</p>
              <p className="image-upload-text">IMAGE</p>
            </label>
            <input
              type="file"
              onChange={(e) => {
                onImageSelection(e);
              }}
              accept="image/*"
              id="img-upload"
            />
          </>
        ) : (
          <img
            src={URL.createObjectURL(img)}
            alt="Selection Display"
            className="uploaded-image"
          />
        )}

        <div className="button-align">
          {buttonActive ? (
            <button className="submit-button-active" type="submit"></button>
          ) : (
            <button className="submit-button" type="submit" disabled></button>
          )}
        </div>
      </form>
    </div>
  );
}

export default CardInput;
