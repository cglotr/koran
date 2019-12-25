import axios from 'axios'

const api = axios.create({
  baseURL: 'https://koran.arikama.co/api'
})

const deleteUserRead = (userId, suraNumber, verseNumber) => {
  return api.delete(`/user/${userId}/read?sura_id=${suraNumber}&verse_id=${verseNumber}`)
}

const getSuraVerse = (suraNumber, verseNumber) => {
  return api.get(`/sura?suraNumber=${suraNumber}&verseNumber=${verseNumber}`)
}

const getSuraVerseTranslation = (suraNumber, verseNumber) => {
  return api.get(`/translation?suraNumber=${suraNumber}&verseNumber=${verseNumber}`)
}

const getUserRead = (userId, suraNumber, verseNumber) => {
  return api.get(`/user/${userId}/read?sura_id=${suraNumber}&verse_id=${verseNumber}`)
}

const postAuth = (idToken) => {
  return api.post('/auth', {
    id_token: idToken
  })
}

const postAuthInvalidate = (id, token) => {
  return api.post(`/auth/${id}/invalidate`, {
    token
  })
}

const postUserRead = (userId, suraNumber, verseNumber) => {
  return api.post(`/user/${userId}/read`, {
    sura_id: parseInt(suraNumber),
    verse_id: parseInt(verseNumber)
  })
}

export default api
export {
  deleteUserRead,
  getSuraVerse,
  getSuraVerseTranslation,
  getUserRead,
  postAuth,
  postAuthInvalidate,
  postUserRead
}
