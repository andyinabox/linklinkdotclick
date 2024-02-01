import { createLink } from './lib/api'

import './main.css'
import './components/link/link'
import './components/evt-btn/evt-btn'

const emitDocumentEvent = (name, detail = {}) => {
  document.dispatchEvent(
    new CustomEvent(name, {
      bubbles: true,
      detail,
    })
  )
}

const startLoading = () => {
  document.body.classList.add('loading')
}
const stopLoading = () => {
  document.body.classList.remove('loading')
}

const linksContainerEl = document.getElementById('links')
const linkElements = linksContainerEl.querySelectorAll('linky-link')
const linkTmpl = document.getElementById('tmpl-link')

const appendLinkElement = (link) => {
  const el = linkTmpl.content.firstElementChild.cloneNode(true)
  linksContainerEl.appendChild(el)
  el.render(link)
}

const deleteLinkElement = (id) => {
  let found = false
  for (const el of linkElements) {
    if (parseInt(el.getAttribute('data-id')) === id) {
      el.remove()
      found = true
      return
    }
  }
  return found
}

const handleCreateLink = async (event) => {
  try {
    startLoading()
    const url = prompt('Enter a website or feed URL')
    const link = await createLink(url)
    appendLinkElement(link)
    emitDocumentEvent('create-link-success', { link })
  } catch (error) {
    console.error(error)
    emitDocumentEvent('create-link-error', { error })
  } finally {
    stopLoading()
  }
}

const handleUpdateLink = async (event) => {
  try {
    startLoading()
    const link = await updateLink(event.detail.link)
    emitDocumentEvent('update-link-success', { link })
  } catch (error) {
    console.error(error)
    emitDocumentEvent('update-link-error', { error })
  } finally {
    stopLoading()
  }
}

const handleDeleteLink = async (event) => {
  try {
    startLoading()
    const { id } = await updateLink(event.detail.link)
    const found = deleteLinkElement(id)
    if (!found) {
      console.warn(
        `link ${id} successfully deleted, but no corresponding DOM element found`
      )
    }
    emitDocumentEvent('delete-link-success', { id })
  } catch (error) {
    console.error(error)
    emitDocumentEvent('delete-link-error', { error })
  } finally {
    stopLoading()
  }
}

document.addEventListener('update-link', (event) => handleUpdateLink(event))
document.addEventListener('create-link', (event) => handleCreateLink(event))
document.addEventListener('delete-link', (event) => handleDeleteLink(event))
