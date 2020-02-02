import _ from 'lodash'
import React, { Component } from 'react'
import { Redirect } from 'react-router-dom'

import { Column } from '@app/components'
import { dimensions, suras } from '@app/constants'

export default class Quran extends Component {
  render () {
    return (
      <Column
        alignItems='center'
        height='200vh'
        justifyContent='flex-start'
        paddingTop={dimensions.PADDING_XXLARGE}
      >
        <Redirect to={`/sura/${this.getRandomSura()}`} />
      </Column>
    )
  }

  getRandomSura = () => {
    const from = 0
    const till = _.keys(suras).length - 1
    return _.random(from, till)
  }
}
