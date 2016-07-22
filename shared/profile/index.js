// @flow
import React, {Component} from 'react'
import {connect} from 'react-redux'
import Render from './render'
import EditProfile from './edit-profile'
import type {Props} from './render'
import flags from '../util/feature-flags'
import {getProfile, updateTrackers} from '../actions/tracker'
import {routeAppend, navigateUp} from '../actions/router'
import {openInKBFS} from '../actions/kbfs'

type State = {loaded: boolean}

class Profile extends Component<void, Props, State> {
  state: State;
  static parseRoute (currentPath, uri) {
    return {
      componentAtTop: {
        title: 'Profile',
        props: {
          username: currentPath.get('username'),
          profileIsRoot: !!uri.count() && uri.last().get('path') === 'root',
        },
      },
      subRoutes: {
        'editprofile': EditProfile,
      },
    }
  }

  componentDidMount () {
    this.props.refresh(this.props.username)
  }

  componentWillReceiveProps (nextProps) {
    if (nextProps.username !== this.props.username) {
      this.props.refresh(nextProps.username)
    }
  }

  constructor () {
    super()
    this.state = {loaded: false}
  }

  render () {
    // TODO: remove this when we actually have loading logic
    setTimeout(() => {
      if (!this.state.loaded) {
        this.setState({loaded: true})
      }
    }, 1.5e3)

    return (
      <Render
        showComingSoon={!flags.tabProfileEnabled}
        {...this.props}
        proofs={this.props.proofs || []}
        loading={!this.state.loaded}
        onBack={!this.props.profileIsRoot ? this.props.onBack : undefined}
        followers={this.props.trackers || []}
        following={this.props.tracking || []}
      />
    )
  }
}

export default connect(
  state => ({
    myUsername: state.config.username,
    trackers: state.tracker.trackers,
  }),
  dispatch => ({
    refresh: username => {
      dispatch(getProfile(username))
      dispatch(updateTrackers(username))
    },
    onUserClick: username => { dispatch(routeAppend({path: 'profile', username})) },
    onBack: () => dispatch(navigateUp()),
    onFolderClick: folder => dispatch(openInKBFS(folder.path)),
    onEditProfile: () => dispatch(routeAppend({path: 'editprofile'})),
  }),
  (stateProps, dispatchProps, ownProps) => {
    const username = ownProps.username || stateProps.myUsername
    const isYou = username === stateProps.myUsername
    const onEditProfile = () => dispatchProps.onEditProfile()
    const bioEditFns = isYou && {
      onBioEdit: onEditProfile,
      onEditAvatarClick: onEditProfile,
      onEditProfile: onEditProfile,
      onLocationEdit: onEditProfile,
      onNameEdit: onEditProfile,
    }

    return {
      ...ownProps,
      ...stateProps.trackers[username],
      ...dispatchProps,
      isYou,
      bioEditFns,
      username,
      refresh: username => dispatchProps.refresh(username),
    }
  }
)(Profile)
