import { call, put, takeEvery, takeLatest } from 'redux-saga/effects'
import api from '@app/api'
import { quran } from '@app/constants'
import { quran as quranSlice } from '@app/slices'

function * fetchSura (action) {
  const suraNumber = parseInt(action.payload)
  const noOfVerses = quran[suraNumber]
  for (let verseNumber = 1; verseNumber <= noOfVerses; verseNumber++) {
    yield put(quranSlice.actions.requestSuraVerse({
      suraNumber,
      verseNumber
    }))
  }
}

function * fetchSuraVerse (action) {
  const suraNumber = action.payload.suraNumber
  const verseNumber = action.payload.verseNumber
  const response = yield call(api.get, `/sura?suraNumber=${suraNumber}&verseNumber=${verseNumber}`)
  const verse = response.data[0]
  yield put(quranSlice.actions.setVerse(verse))
}

export default function * () {
  yield takeEvery(quranSlice.actions.requestSuraVerse, fetchSuraVerse)
  yield takeLatest(quranSlice.actions.requestSura, fetchSura)
}
