import produce from 'immer'
import _ from 'lodash'
import { createSlice } from '@reduxjs/toolkit'

export default createSlice({
  initialState: {},
  name: 'quran',
  reducers: {
    requestSura: (state) => state,
    requestSuraVerse: (state) => state,
    setVerse: (state, action) => produce(state, draft => {
      const ayah = action.payload.ayah
      const suraNumber = action.payload.sura_number.toString()
      const verseNumber = action.payload.verse_number.toString()
      _.merge(draft, {
        [suraNumber]: {
          [verseNumber]: {
            ayah
          }
        }
      })
    })
  }
})
