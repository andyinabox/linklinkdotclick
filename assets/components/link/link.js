import { getLink, updateLink, deleteLink } from '../../lib/api'
import { handleError } from '../../lib/errors'
import { Component } from '../component'
export class Link extends Component {
  constructor() {
    super()
  }

  async fetchData() {
    const id = this.getAttribute('data-id')
    if (!id) return
    try {
      this.loading = true
      this.data = await getLink(id)
    } catch (err) {
      handleError(err)
    } finally {
      this.loading = false
    }
  }

  // render data to the element
  render() {
    const data = this.data
    this.setAttribute('data-id', this.data.id)
    this.slots.link.href = data.siteUrl
    this.slots.link.textContent = data.siteName

    if (data.hideUnreadCount) {
      this.slots.count.textContent = ''
    } else if (data.unreadCount) {
      this.slots.count.textContent = data.unreadCount
    } else if (!data.feedUrl) {
      this.slots.count.textContent = '?'
    }

    this.slots['edit-menu'].querySelector('[name="site-url"]').value =
      data.siteUrl
    this.slots['edit-menu'].querySelector('[name="feed-url"]').value =
      data.feedUrl
    this.setAttribute('data-id', data.id)
  }

  async handleClick() {
    const link = this.data
    link.lastClicked = new Date().toJSON()
    try {
      const updatedLink = await updateLink(link)
      this.data = updatedLink
      this.dispatchEvent(new CustomEvent('link-click', { bubbles: true }))
    } catch (err) {
      handleError(err)
    }
  }

  async handleDelete(e) {
    const { id, siteName } = this.data

    if (!confirm(`Delete link "${siteName}"?`)) return

    try {
      this.loading = true
      await deleteLink(id)
      this.remove()
    } catch (err) {
      handleError(err)
    } finally {
      this.loading = false
    }
  }

  async handleFieldEdit(fieldName, promptText) {
    const data = this.data
    const value = prompt(promptText)
    if (!value) return
    data[fieldName] = value
    try {
      this.loading = true
      const updated = await updateLink(data)
      this.data = updated
    } catch (err) {
      handleError(err)
    } finally {
      this.loading = false
    }
  }
  async handleHideUnreadCount(evt) {
    const data = this.data
    const isChecked = evt.target.checked
    data.hideUnreadCount = isChecked
    console.log('handleHideUnreadCount', isChecked)
    try {
      this.loading = true
      const updated = await updateLink(data)
      this.data = updated
    } catch (err) {
      handleError(err)
    } finally {
      this.loading = false
    }
  }

  connectedCallback() {
    const data = this.data
    const edit = this.slots['edit-menu']
    const deleteBtn = edit.querySelector('[name="delete"]')
    const renameBtn = edit.querySelector('[name="rename"]')
    const editSiteUrlButton = edit.querySelector('[name="edit-site-url"]')
    const editFeedUrlButton = edit.querySelector('[name="edit-feed-url"]')
    const hideUnreadCountCheckbox = edit.querySelector(
      '[name="hide-unread-count"]'
    )

    this.listen(edit, 'submit', (e) => e.preventDefault())
    this.listen(deleteBtn, 'click', this.handleDelete)
    this.listen(renameBtn, 'click', () =>
      this.handleFieldEdit(
        'siteName',
        `Enter a new name for "${data.siteName}"`
      )
    )
    this.listen(editSiteUrlButton, 'click', () =>
      this.handleFieldEdit('siteUrl', `Enter a new site url`)
    )
    this.listen(editFeedUrlButton, 'click', () =>
      this.handleFieldEdit('feedUrl', `Enter a new RSS/Atom feed url`)
    )
    this.listen(hideUnreadCountCheckbox, 'change', this.handleHideUnreadCount)
    this.listen(this.slots.link, 'click', this.handleClick)
  }
}

Link.create = function (parentEl, data) {
  const linkTmpl = document.getElementById('tmpl-link')
  const linkEl = linkTmpl.content.firstElementChild.cloneNode(true)
  parentEl.prepend(linkEl)
  linkEl.data = data
  return linkEl
}

customElements.define('linky-link', Link)
