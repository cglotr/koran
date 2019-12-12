import { call, put, takeEvery, takeLatest } from 'redux-saga/effects'
import api from '@app/api'
import { quran } from '@app/constants'
import { quran as quranSlice } from '@app/slices'

function * fetchSura (action) {
  const suraNumber = parseInt(action.payload)
  const noOfVerses = quran[suraNumber]
  for (let verseNumber = 1; verseNumber <= noOfVerses; verseNumber++) {
    yield put(quranSlice.actions.fetchSuraVerse({
      suraNumber,
      verseNumber
    }))
  }
}

function * fetchSuraVerse (action) {
  const suraNumber = action.payload.suraNumber
  const verseNumber = action.payload.verseNumber
  yield call(api.get, `/sura?suraNumber=${suraNumber}&verseNumber=${verseNumber}`)
}

export default function * () {
  yield takeEvery(quranSlice.actions.fetchSuraVerse, fetchSuraVerse)
  yield takeLatest(quranSlice.actions.fetchSura, fetchSura)
}
