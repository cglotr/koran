import firebase from 'firebase/app'
import { call, put, takeLatest } from 'redux-saga/effects'

import {
  app as appSlice,
  user as userSlice
} from '@app/slices'

const auth = firebase.auth()
const provider = new firebase.auth.GoogleAuthProvider()

function * fetchSignIn () {
  const response = yield call([auth, auth.signInWithPopup], provider)
  yield put(userSlice.actions.setUserEmail(response.user.email))
}

export default function * () {
  yield takeLatest(appSlice.actions.requestSignIn, fetchSignIn)
}
