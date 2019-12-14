import _ from 'lodash'
import { call, delay, put, takeEvery, takeLatest } from 'redux-saga/effects'
import { getSuraVerse, getSuraVerseTranslation } from '@app/api'
import { durations, suras } from '@app/constants'
import { quran as quranSlice } from '@app/slices'

function * fetchSura (action) {
  const suraNumber = parseInt(action.payload)
  const noOfVerses = _.get(suras, [suraNumber, 'numberOfVerses'])
  for (let verseNumber = 1; verseNumber <= noOfVerses; verseNumber++) {
    yield put(quranSlice.actions.requestSuraVerse({
      suraNumber,
      verseNumber
    }))
    yield delay(durations.SURA_LOAD_DELAY)
  }
}

function * fetchSuraVerse (action) {
  const suraNumber = action.payload.suraNumber
  const verseNumber = action.payload.verseNumber
  const response = yield call(getSuraVerse, suraNumber, verseNumber)
  const verse = response.data[0]
  yield put(quranSlice.actions.setVerse(verse))
}

function * fetchSuraVerseTranslation (action) {
  const suraNumber = action.payload.suraNumber
  const verseNumber = action.payload.verseNumber
  const response = yield call(getSuraVerseTranslation, suraNumber, verseNumber)
  const translation = response.data[0]
  yield put(quranSlice.actions.setVerseTranslation(translation))
}

export default function * () {
  yield takeEvery(quranSlice.actions.requestSuraVerse, fetchSuraVerse)
  yield takeEvery(quranSlice.actions.requestSuraVerse, fetchSuraVerseTranslation)
  yield takeLatest(quranSlice.actions.requestSura, fetchSura)
}
