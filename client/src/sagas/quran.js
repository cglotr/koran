import { call, takeLatest } from 'redux-saga/effects'
import { quran } from '@app/slices'

function * fetchSura () {
  yield call(fetch, 'http://koran.cglotr.com/api/sura?suraNumber=1&verseNumber=1')
}

export default function * () {
  yield takeLatest(quran.actions.fetchSura, fetchSura)
}
