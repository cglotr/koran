import firebase from 'firebase/app'
import 'firebase/auth'
import { call, put, select, takeLatest } from 'redux-saga/effects'

import { postAuth, postAuthInvalidate } from '@app/api'
import { keys } from '@app/constants'
import {
  app as appSlice,
  read as readSlice,
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
  const { ma } = response.user
  const authResponse = yield call(postAuth, ma)
  const { id, uid, token } = authResponse.data
  yield put(userSlice.actions.setUser({ id, uid, token }))
  window.localStorage.setItem(keys.PERSISTED_USER_STATE, JSON.stringify({
    id,
    token,
    uid
  }))
}

function * fetchSignOut () {
  const userId = yield select((state) => state.user.id)
  const userToken = yield select((state) => state.user.token)
  try {
    yield call(postAuthInvalidate, userId, userToken)
  } catch (e) {
    console.error(e)
  }
  yield put(userSlice.actions.reset())
  yield put(readSlice.actions.reset())
  yield put(appSlice.actions.setIsUserDialogOpen(false))
  window.localStorage.setItem(keys.PERSISTED_USER_STATE, '')
}

export default function * () {
  yield takeLatest(appSlice.actions.requestSignIn, fetchSignIn)
  yield takeLatest(userSlice.actions.requestSignOut, fetchSignOut)
}
