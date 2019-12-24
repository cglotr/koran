import firebase from 'firebase/app'
import 'firebase/auth'
import { call, put, takeLatest } from 'redux-saga/effects'

import {
  app as appSlice,
  user as userSlice
} from '@app/slices'

firebase.initializeApp({
  apiKey: process.env.CLIENT_FIREBASE_API_KEY,
  appId: '1:223159991719:web:c5e2155b09698732eaca82',
  authDomain: 'koran-259911.firebaseapp.com',
  databaseURL: 'https://koran-259911.firebaseio.com',
  measurementId: 'G-BPJRSG1QRB',
  messagingSenderId: '223159991719',
  projectId: 'koran-259911',
  storageBucket: 'koran-259911.appspot.com'
})

const auth = firebase.auth()
const provider = new firebase.auth.GoogleAuthProvider()

function * fetchSignIn () {
  const response = yield call([auth, auth.signInWithPopup], provider)
  const { email } = response.user
  yield put(userSlice.actions.setUserEmail(email))
}

export default function * () {
  yield takeLatest(appSlice.actions.requestSignIn, fetchSignIn)
}
