import React, { Component } from 'react'
import { Typography } from '@material-ui/core'

import { Column } from '@app/components'

export default class Quran extends Component {
  render () {
    return (
      <Column alignItems='center' justifyContent='center'>
        <Typography gutterBottom variant='h2'>Koran</Typography>
        <Typography align='center' variant='subtitle1'>
          Koran is still in the Beta development phase &amp; we are actively developing it. We have a few features in mind that we think will make Koran the best place to read the Quran on the internet. And we are very excited about this mission!
        </Typography>
      </Column>
    )
  }
}
