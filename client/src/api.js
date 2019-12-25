import axios from 'axios'

const api = axios.create({
  baseURL: 'https://koran.arikama.co/api'
})

const getSuraVerse = (suraNumber, verseNumber) => {
  return api.get(`/sura?suraNumber=${suraNumber}&verseNumber=${verseNumber}`)
}

const getSuraVerseTranslation = (suraNumber, verseNumber) => {
  return api.get(`/translation?suraNumber=${suraNumber}&verseNumber=${verseNumber}`)
}

const postAuth = (idToken) => {
  return api.post(`/auth`, {
    id_token: idToken
  })
}

export default api
export {
  getSuraVerse,
  getSuraVerseTranslation,
  postAuth
}
