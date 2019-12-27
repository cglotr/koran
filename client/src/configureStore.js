import createSagaMiddleware from 'redux-saga'
import { configureStore } from '@reduxjs/toolkit'

import { keys } from '@app/constants'
import {
  quran as quranSaga,
  read as readSaga,
  user as userSaga
} from '@app/sagas'
import {
  app as appSlice,
  quran as quranSlice,
  read as readSlice,
  user as userSlice
} from '@app/slices'
import { getPersistedUserState } from '@app/utils'

const preloadedState = {
  user: getPersistedUserState(keys.PERSISTED_USER_STATE)
}
const sagaMiddleware = createSagaMiddleware()

export default () => {
  const store = configureStore({
    middleware: [
      sagaMiddleware
    ],
    preloadedState,
    reducer: {
      [appSlice.name]: appSlice.reducer,
      [quranSlice.name]: quranSlice.reducer,
      [readSlice.name]: readSlice.reducer,
      [userSlice.name]: userSlice.reducer
    }
  })
  sagaMiddleware.run(quranSaga)
  sagaMiddleware.run(readSaga)
  sagaMiddleware.run(userSaga)
  return store
}
