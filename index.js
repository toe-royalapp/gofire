// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
import { getAnalytics } from "firebase/analytics";
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
  apiKey: "AIzaSyCwXHvwNSJKyWaPUa58hNpu5ny7s-C7cG4",
  authDomain: "gofire-9ce2c.firebaseapp.com",
  projectId: "gofire-9ce2c",
  storageBucket: "gofire-9ce2c.appspot.com",
  messagingSenderId: "646521782984",
  appId: "1:646521782984:web:e1f295dfa8e4afdb429bfd",
  measurementId: "G-WMJ0H2TK22"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
const analytics = getAnalytics(app);