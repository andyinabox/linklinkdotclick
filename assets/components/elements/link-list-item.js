import { eventsMixin, slotsMixin } from '../../lib/mixins'
import { updateLink, deleteLink, clickLink, getLink } from '../../lib/api'
import { handleError } from '../../lib/errors'
class Link extends HTMLLIElement {
  constructor() {
    super()
    this.registerSlots()

    // replace link href
    const href = this.slots.link.getAttribute('data-href')
    if (href) {
      this.siteUrl = href
      this.removeAttribute('data-href')
    }

    // disable form submit
    this.form = this.slots.form
    this.form.onsubmit = (evt) => evt.preventDefault()

    // collect form inputs and buttons for convenience
    this.inputs = {}
    this.form.querySelectorAll('input[name]').forEach((el) => {
      this.inputs[el.getAttribute('name')] = el
    })
    this.buttons = {}
    this.form.querySelectorAll('button[name]').forEach((el) => {
      this.buttons[el.getAttribute('name')] = el
    })

    // disable save button by default
    this.buttons['link-item-save'].disabled = true

    // fetch data if there is a link id
    if (this.linkId) {
      this.fetchData()
    }
  }
  async fetchData() {
    try {
      this.loading = true
      const link = await getLink(this.linkId)
      this.data = link
    } catch (err) {
      handleError(err)
    } finally {
      this.loading = false
    }
  }
  get data() {
    return { ...this._data }
  }
  set data(d) {
    this._data = d
    this.render()
  }
  get loading() {
    return this.classList.hasClass('loading')
  }
  set loading(v) {
    if (v) {
      this.classList.add('loading')
      this.broadcast('loading-start')
    } else {
      this.classList.remove('loading')
      this.broadcast('loading-stop')
    }
  }

  set linkId(v) {
    this.setAttribute('data-id', v)
  }
  get linkId() {
    return parseInt(this.getAttribute('data-id'))
  }
  set siteUrl(v) {
    this.slots.link.href = v
  }
  get siteUrl() {
    return this.slots.link.href
  }
  set siteName(v) {
    this.slots.link.innerText = v
  }
  get siteName() {
    return this.slots.link.innerText
  }
  set unreadCount(v) {
    this.slots.count.innerText = v
  }
  get unreadCount() {
    return parseInt(this.slots.count.innerText)
  }

  set formLinkId(v) {
    this.inputs['id'].value = v
  }
  get formLinkId() {
    return parseInt(this.inputs['id'].value)
  }

  set formSiteName(v) {
    this.inputs['site-name'].value = v
  }
  get formSiteName() {
    return this.inputs['site-name'].value
  }

  set formSiteUrl(v) {
    this.inputs['site-url'].value = v
  }
  get formSiteUrl() {
    return this.inputs['site-url'].value
  }

  set formFeedUrl(v) {
    this.inputs['feed-url'].value = v
  }
  get formFeedUrl() {
    return this.inputs['feed-url'].value
  }

  set formHideUnreadCount(v) {
    this.inputs['hide-unread-count'].checked = v
  }
  get formHideUnreadCount() {
    return this.inputs['hide-unread-count'].checked
  }

  set formLastClicked(v) {
    this.inputs['last-clicked'].value = v
  }
  get formLastClicked() {
    return this.inputs['last-clicked'].value
  }

  get formData() {
    return {
      id: this.formLinkId,
      lastClicked: this.formLastClicked,
      siteName: this.formSiteName,
      siteUrl: this.formSiteUrl,
      feedUrl: this.formFeedUrl,
      hideUnreadCount: this.formHideUnreadCount,
    }
  }

  render() {
    const data = this.data
    if (!data) return
    const { id, siteName, unreadCount, siteUrl, hideUnreadCount, feedUrl } =
      data

    if (hideUnreadCount) {
      this.slots.count.textContent = ''
    } else if (unreadCount) {
      this.slots.count.textContent = data.unreadCount
    } else if (!feedUrl) {
      this.slots.count.textContent = '?'
    } else {
      this.slots.count.textContent = ''
    }

    this.linkId = id
    this.siteName = siteName
    this.siteUrl = siteUrl
    this.renderForm()
  }

  handleFormInput() {
    const { id, siteName, siteUrl, feedUrl, hideUnreadCount } = this.data
    const changed =
      this.formLinkId !== id ||
      this.formSiteName !== siteName ||
      this.formSiteUrl !== siteUrl ||
      this.formFeedUrl !== feedUrl ||
      this.formHideUnreadCount !== hideUnreadCount

    this.buttons['link-item-save'].disabled = !changed
  }

  renderForm() {
    const data = this.data
    const { id, siteName, siteUrl, feedUrl, hideUnreadCount, lastClicked } =
      data
    this.formLinkId = id
    this.formLastClicked = lastClicked
    this.formSiteName = siteName
    this.formSiteUrl = siteUrl
    this.formFeedUrl = feedUrl
    this.formHideUnreadCount = hideUnreadCount
    this.handleFormInput()
  }

  async onClick() {
    try {
      this.loading = true
      const link = await clickLink(this.linkId)
      this.data = link
      this.broadcast('link-click-success')
    } catch (err) {
      handleError(err)
    } finally {
      this.loading = false
    }
  }
  async onUpdate() {
    try {
      this.loading = true
      const link = await updateLink(Object.assign(this.data, this.formData))
      this.data = link
      this.broadcast('link-update-success')
    } catch (err) {
      handleError(err)
    } finally {
      this.loading = false
    }
  }
  async onDelete() {
    if (!confirm(`Are you sure you want to delete ${this.data.siteName}?`))
      return
    try {
      this.loading = true
      const result = await deleteLink(this.data.id)
      this.broadcast('link-delete-success')
      this.remove()
    } catch (err) {
      handleError(err)
    } finally {
      this.loading = false
    }
  }

  connectedCallback() {
    this.listen(this.slots.link, 'click', this.onClick)
    this.listen(this.buttons['link-item-delete'], 'click', this.onDelete)
    this.listen(this.buttons['link-item-save'], 'click', this.onUpdate)
    Object.values(this.inputs).forEach((el) =>
      this.listen(el, 'input', this.handleFormInput)
    )
  }
  disconnectedCallback() {
    this.unlistenAll()
  }
}
Object.assign(Link.prototype, eventsMixin, slotsMixin)
customElements.define('link-list-item', Link, { extends: 'li' })
