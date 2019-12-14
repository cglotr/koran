import axios from 'axios'

const api = axios.create({
  baseURL: 'http://koran.cglotr.com/api'
})

const getSuraVerse = (suraNumber, verseNumber) => {
  return api.get(`/sura?suraNumber=${suraNumber}&verseNumber=${verseNumber}`)
}

export default api
export {
  getSuraVerse
}
