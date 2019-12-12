import createSagaMiddleware from 'redux-saga'
import { configureStore } from '@reduxjs/toolkit'
import { quran as quranSaga } from '@app/sagas'
import { quran as quranSlice } from '@app/slices'

const sagaMiddleware = createSagaMiddleware()

export default () => {
  const store = configureStore({
    middleware: [
      sagaMiddleware
    ],
    reducer: {
      [quranSlice.name]: quranSlice.reducer
    }
  })
  sagaMiddleware.run(quranSaga)
  return store
}
