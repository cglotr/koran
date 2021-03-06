import { call, put, select, takeEvery } from 'redux-saga/effects'

import { deleteUserRead, getUserRead, postUserRead } from '@app/api'
import { user as userSelector } from '@app/selectors'
import { quran as quranSlice, read as readSlice } from '@app/slices'

function * fetchUserRead (action) {
  const isUserSignedIn = yield select(userSelector.isSignedIn)
  if (!isUserSignedIn) {
    return
  }
  const userId = yield select(userSelector.id)
  const { suraNumber, verseNumber } = action.payload
  const response = yield call(getUserRead, userId, suraNumber, verseNumber)
  const isRead = response.data.read
  yield put(readSlice.actions.setSuraVerseRead({
    isRead,
    suraNumber,
    verseNumber
  }))
}

function * updateUserRead (action) {
  const isUserSignedIn = yield select(userSelector.isSignedIn)
  if (!isUserSignedIn) {
    return
  }
  const userId = yield select(userSelector.id)
  const { isRead, suraNumber, verseNumber } = action.payload
  if (isRead) {
    yield call(postUserRead, userId, suraNumber, verseNumber)
  } else {
    yield call(deleteUserRead, userId, suraNumber, verseNumber)
  }
}

export default function * () {
  yield takeEvery(quranSlice.actions.requestSuraVerse, fetchUserRead)
  yield takeEvery(readSlice.actions.setSuraVerseRead, updateUserRead)
}
