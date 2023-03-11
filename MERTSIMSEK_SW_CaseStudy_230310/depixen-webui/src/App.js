import Axios from "axios";
import "./App.css";
import { useState, useEffect } from "react";
import CardInput from "./components/cardInput/CardInput";
import CardView from "./components/cardView/CardView";
import { storage } from "./firebase";
import { ref, uploadBytesResumable, getDownloadURL } from "firebase/storage";
import { v4 } from "uuid";

function App() {
  const [title, setTitle] = useState("New title");
  const [desc, setDesc] = useState("New description");
  const [img, setImg] = useState(null);
  const [buttonActive, setButtonActive] = useState(false);
  const [imgDownloadURL, setImgDownloadURL] = useState("");
  const [uploadedCard, setUploadedCard] = useState({
    title: "",
    description: "",
    imageuri: "",
    createddate: "",
  });

  useEffect(() => {
    if (imgDownloadURL !== "") {
      postCard();
    }
  }, [imgDownloadURL]);

  const createDateNowForPostgres = () => {
    const now = new Date();
    const year = now.getFullYear();
    const month = String(now.getMonth() + 1).padStart(2, "0");
    const day = String(now.getDate()).padStart(2, "0");
    const hours = String(now.getHours()).padStart(2, "0");
    const minutes = String(now.getMinutes()).padStart(2, "0");
    const seconds = String(now.getSeconds()).padStart(2, "0");

    return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
  };

  const handleClick = (e) => {
    e.target.select();
  };

  const resetCardInputStates = () => {
    setTitle("New title");
    setDesc("New description");
    setImg(null);
    setButtonActive(false);
  };

  const postCard = () => {
    const card = {
      title: title,
      description: desc,
      imageuri: imgDownloadURL,
      createddate: createDateNowForPostgres(),
    };

    Axios.post("http://localhost:8080/cards", card)
      .then((response) => {
        console.log("Post request success:", response.data);
        setUploadedCard(response.data.card);
        resetCardInputStates();
      })
      .catch((error) => {
        console.error("Post request error:", error);
      });
  };

  const onImageSelection = (e) => {
    const imageFile = e.target.files[0];
    if (imageFile && imageFile.type.startsWith("image/")) {
      setImg(imageFile);
      setButtonActive(true);
    } else {
      setImg(null);
      alert("Please select a valid image file.");
    }
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    handleImageUpload();
  };

  const handleImageUpload = () => {
    if (img == null) {
      alert("An Error Occured during upload. Please try again.");
      return;
    } else {
      const imageRef = ref(storage, `images/${img.name + v4()}`);
      const uploadTask = uploadBytesResumable(imageRef, img);

      uploadTask.on("state_changed", null, null, () => {
        getDownloadURL(uploadTask.snapshot.ref).then((downloadURL) => {
          setImgDownloadURL(downloadURL);
          console.log(imgDownloadURL);
        });
      });
    }
  };

  return (
    <div className="page">
      <CardInput
        handleClick={handleClick}
        title={title}
        setTitle={setTitle}
        desc={desc}
        setDesc={setDesc}
        img={img}
        setImg={setImg}
        onImageSelection={onImageSelection}
        buttonActive={buttonActive}
        handleSubmit={handleSubmit}
      />
      <CardView imgDownloadURL={imgDownloadURL} uploadedCard={uploadedCard} />
    </div>
  );
}

export default App;
