import { initializeApp } from "firebase/app";
import { getStorage } from "firebase/storage"
// Your web app's Firebase configuration
const firebaseConfig = {
  apiKey: "AIzaSyB02dz7cTVROHGrZu1Vh99NJf0iOhUh_WI",
  authDomain: "depixensw.firebaseapp.com",
  projectId: "depixensw",
  storageBucket: "depixensw.appspot.com",
  messagingSenderId: "1056976244613",
  appId: "1:1056976244613:web:97f586aeac075162493eec"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);

export const storage = getStorage(app);
