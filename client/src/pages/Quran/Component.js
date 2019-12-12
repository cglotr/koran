import _ from 'lodash'
import React, { Component } from 'react'
import { Link } from 'react-router-dom'
import { quran } from '@app/constants'

export default class Quran extends Component {
  render () {
    return (
      <>
        <h1>Quran</h1>
        <div>
          {
            _.keys(quran)
              .map((key) => parseInt(key))
              .sort((a, b) => a - b)
              .map((suraNumber) => {
                return (
                  <h2 key={suraNumber}>
                    <Link to={`/sura/${suraNumber}`}>{suraNumber}</Link>
                  </h2>
                )
              })
          }
        </div>
      </>
    )
  }
}
