importScripts(
  "https://www.gstatic.com/firebasejs/10.7.1/firebase-app-compat.js"
);
importScripts(
  "https://www.gstatic.com/firebasejs/10.7.1/firebase-messaging-compat.js"
);

firebase.initializeApp({
  apiKey: "AIzaSyBT2zE5VRCiie6bsA2kEZBAkaieECpeLWM",
  authDomain: "event-driven-go.firebaseapp.com",
  projectId: "event-driven-go",
  storageBucket: "event-driven-go.firebasestorage.app",
  messagingSenderId: "100381994257",
  appId: "1:100381994257:web:8c2667e42ae12cd4f0b0b0",
  measurementId: "G-THJV5Z6CC3",
});

const messaging = firebase.messaging()

messaging.onBackgroundMessage((payload) => {
  console.log(
    "[firebase-messaging-sw.js] Recibido mensaje en segundo plano:",
    payload
  );
  self.registration.showNotification(payload.notification.title, {
    body: payload.notification.body,
    icon: "/assets/icons/icon-192x192.png",
  });
});
