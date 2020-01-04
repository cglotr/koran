import _ from 'lodash'
import PropTypes from 'prop-types'
import React from 'react'
import { Column, Footer, Verse } from '@app/components'
import { Typography } from '@material-ui/core'
import {
  LocalLibraryOutlined as LocalLibraryOutlinedIcon,
  PlaylistAddCheck as PlaylistAddCheckIcon
} from '@material-ui/icons'

import { dimensions, suras } from '@app/constants'
import Fab from './Fab'
import LinearProgress from './LinearProgress'
import PaddedColumn from './PaddedColumn'
import VerseTypography from './VerseTypography'

export default class Component extends React.Component {
  static propTypes = {
    isUserSignedIn: PropTypes.bool.isRequired,
    match: PropTypes.shape({
      params: PropTypes.shape({
        number: PropTypes.string.isRequired
      })
    }),
    quran: PropTypes.object.isRequired,
    read: PropTypes.object,
    requestSura: PropTypes.func.isRequired,
    setSuraVerseRead: PropTypes.func.isRequired
  }

  state = {
    isShowingOnlyUnread: false
  }

  componentDidMount () {
    this.props.requestSura()
  }

  componentDidUpdate (prevProps) {
    const prevSuraNumber = _.get(prevProps, ['match', 'params', 'number'])
    const suraNumber = _.get(this.props, ['match', 'params', 'number'])
    if (suraNumber !== prevSuraNumber || this.props.isUserSignedIn !== prevProps.isUserSignedIn) {
      this.props.requestSura()
    }
  }

  render () {
    const suraNumber = _.get(this.props, ['match', 'params', 'number'])
    const progress = this.getProgress(suraNumber)
    return (
      <Column>
        {this.renderInfo(progress)}
        {this.renderVerses(suraNumber)}
        <Footer />
        {this.renderFab()}
      </Column>
    )
  }

  renderFab = () => {
    if (!this.props.isUserSignedIn) {
      return null
    }
    return (
      <Fab color='default' onClick={this.handleFabClick}>
        {this.state.isShowingOnlyUnread ? <PlaylistAddCheckIcon /> : <LocalLibraryOutlinedIcon />}
      </Fab>
    )
  }

  renderInfo = (progress) => {
    if (this.props.isUserSignedIn && this.state.isShowingOnlyUnread) {
      return (
        <Column paddingTop={dimensions.PADDING_LARGE} />
      )
    }
    const suraNumber = this.props.match.params.number
    const suraName = _.get(suras, [suraNumber, 'suraName'])
    const suraNameTranslation = _.get(suras, [suraNumber, 'suraNameTranslation'])
    return (
      <Column>
        <PaddedColumn>
          <Typography align='center' variant='h5'>{suraName}</Typography>
          <Typography align='center' variant='subtitle1'>{suraNameTranslation}</Typography>
        </PaddedColumn>
        <LinearProgress value={progress} variant='determinate'/>
        {this.renderBismillah()}
      </Column>
    )
  }

  renderVerses = (suraNumber) => {
    const sura = _.get(this.props.quran, suraNumber)
    const verses = _.keys(sura).sort((a, b) => parseInt(a) - parseInt(b))
    return verses.map((verse) => {
      let ayah = _.get(sura, [verse, 'ayah'])
      if (ayah && parseInt(suraNumber) > 1 && parseInt(verse) === 1) {
        ayah = ayah.replace(/^(بِسْمِ اللَّهِ الرَّحْمَٰنِ الرَّحِيمِ)/, '')
      }
      const isRead = _.get(this.props.read, [suraNumber, verse], false)
      const translation = _.get(sura, [verse, 'translation'])
      if (this.props.isUserSignedIn && this.state.isShowingOnlyUnread && isRead) {
        return null
      }
      return (
        <Verse
          ayah={ayah}
          isCheckboxEnabled={this.props.isUserSignedIn}
          isRead={isRead}
          key={verse}
          onCheckboxChange={this.handleCheckboxChange(verse)}
          translation={translation}
          verseNumber={verse}
        />
      )
    })
  }

  getProgress = (suraNumber) => {
    const sura = _.get(this.props, ['quran', suraNumber], {})
    const loaded = _.keys(sura).length
    const total = _.get(suras, [suraNumber, 'numberOfVerses'])
    return _.round(100 * (loaded / total))
  }

  handleCheckboxChange = (verseNumber) => (isRead) => {
    const suraNumber = this.props.match.params.number
    this.props.setSuraVerseRead(suraNumber, verseNumber, isRead)
  }

  handleFabClick = () => {
    this.setState({ isShowingOnlyUnread: !this.state.isShowingOnlyUnread })
  }

  renderBismillah = () => {
    if (parseInt(this.props.match.params.number) === 1) return null
    return (
      <Column>
        <VerseTypography align='center' gutterBottom variant='h3'>بِسْمِ اللَّهِ الرَّحْمَٰنِ الرَّحِيمِ</VerseTypography>
        <VerseTypography
          align='center'
          gutterBottom
        >
          In the name of Allah, Most Gracious, Most Merciful.
        </VerseTypography>
      </Column>
    )
  }
}
