import firebase from 'firebase/app'
import { takeLatest } from 'redux-saga/effects'

import { app as appSlice } from '@app/slices'

const provider = new firebase.auth.GoogleAuthProvider()

function * fetchSignIn (action) {
  firebase.auth().signInWithPopup(provider)
    .then((result) => {
      console.log('result', result)
    })
    .catch((error) => {
      console.log('error', error)
    })
}

export default function * () {
  yield takeLatest(appSlice.actions.requestSignIn, fetchSignIn)
}
