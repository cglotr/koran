import createSagaMiddleware from 'redux-saga'
import { configureStore } from '@reduxjs/toolkit'
import { quran as quranSaga, user as userSaga } from '@app/sagas'
import { app as appSlice, quran as quranSlice } from '@app/slices'

const sagaMiddleware = createSagaMiddleware()

export default () => {
  const store = configureStore({
    middleware: [
      sagaMiddleware
    ],
    reducer: {
      [appSlice.name]: appSlice.reducer,
      [quranSlice.name]: quranSlice.reducer
    }
  })
  sagaMiddleware.run(quranSaga)
  sagaMiddleware.run(userSaga)
  return store
}
