import React from 'react'
import './CardView.css'

function CardView({imgDownloadURL, uploadedCard}) {
  return (
    <div>
        <h1 className="title-above-card">New Title</h1>
        <div className="card-with-data">
        <input type="text" wrap="soft" value={uploadedCard.title} className="title-in-card" disabled/>
          <textarea type="text" value={uploadedCard.description} className="description-in-card" disabled/>
          <div className="main-image-background">
          { imgDownloadURL && <img
            className="uploaded-image"
            src={imgDownloadURL}
            alt=""
          />}
          </div>
        </div>
      </div>
  )
}

export default CardView