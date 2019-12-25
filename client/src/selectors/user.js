import _ from 'lodash'

const isSignedIn = (state) => {
  const userId = _.get(state, ['user', 'id'])
  const userToken = _.get(state, ['user', 'token'])
  return userId && userToken
}

export default {
  isSignedIn
}
