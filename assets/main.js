import './components/elements/form-add-link'
import './components/elements/form-edit-mode'
import './components/elements/form-style-editor'
import './components/elements/link-list-item'
import './components/elements/link-list'

import './components/layouts/page-home'

// remove query
if (window.location.search) {
  const split = window.location.toString().split('?')
  history.replaceState({}, '', split[0])
}
