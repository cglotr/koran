export default (key) => {
  try {
    const value = window.localStorage.getItem(key)
    if (!value) {
      return {}
    }
    return JSON.parse(value)
  } catch {
    return {}
  }
}
