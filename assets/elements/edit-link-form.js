import { handleError } from '../lib/errors'
import { deleteLink, updateLink } from '../lib/api'
import { renderDataMixin } from '../lib/mixins'
import { ProgressiveForm } from './progressive-form'
export class EditLinkForm extends ProgressiveForm {
  constructor() {
    super()
    this.buttons['link-item-save'].disabled = true
    // todo: this will need to be handled differently when creting new links
    this.data = this.formData
  }

  get formData() {
    return {
      id: this.linkId,
      lastClicked: this.lastClicked,
      siteName: this.siteName,
      siteUrl: this.siteUrl,
      feedUrl: this.feedUrl,
      hideUnreadCount: this.hideUnreadCount,
    }
  }

  set linkId(v) {
    this.inputs['id'].value = v
  }
  get linkId() {
    return parseInt(this.inputs['id'].value)
  }

  set siteName(v) {
    this.inputs['site-name'].value = v
  }
  get siteName() {
    return this.inputs['site-name'].value
  }

  set siteUrl(v) {
    this.inputs['site-url'].value = v
  }
  get siteUrl() {
    return this.inputs['site-url'].value
  }

  set feedUrl(v) {
    this.inputs['feed-url'].value = v
  }
  get feedUrl() {
    return this.inputs['feed-url'].value
  }

  set hideUnreadCount(v) {
    this.inputs['hide-unread-count'].checked = v
  }
  get hideUnreadCount() {
    return this.inputs['hide-unread-count'].checked
  }

  set lastClicked(v) {
    this.inputs['last-clicked'].value = v
  }
  get lastClicked() {
    return this.inputs['last-clicked'].value
  }

  handleInput() {
    const { id, siteName, siteUrl, feedUrl, hideUnreadCount } = this.data
    const changed =
      this.linkId !== id ||
      this.siteName !== siteName ||
      this.siteUrl !== siteUrl ||
      this.feedUrl !== feedUrl ||
      this.hideUnreadCount !== hideUnreadCount

    this.buttons['link-item-save'].disabled = !changed
  }
  async save() {
    try {
      this.broadcast('loading-start')
      const link = await updateLink(this.formData)
      this.data = link
      this.broadcast('update-link-success', { link })
    } catch (err) {
      handleError(err)
    } finally {
      this.broadcast('loading-end')
    }
  }
  async delete() {
    if (!confirm(`Are you sure you want to delete ${this.data.siteName}?`))
      return
    try {
      this.broadcast('loading-start')
      const result = await deleteLink(this.formData.id)
      this.broadcast('delete-link-success', result)
    } catch (err) {
      handleError(err)
    } finally {
      this.broadcast('loading-end')
    }
  }
  render() {
    const { id, siteName, siteUrl, feedUrl, hideUnreadCount, lastClicked } =
      this.data
    this.linkId = id
    this.lastClicked = lastClicked
    this.siteName = siteName
    this.siteUrl = siteUrl
    this.feedUrl = feedUrl
    this.hideUnreadCount = hideUnreadCount
    this.handleInput()
  }
  connectedCallback() {
    this.listen(this.buttons['link-item-delete'], 'click', this.delete)
    this.listen(this.buttons['link-item-save'], 'click', this.save)
    Object.values(this.inputs).forEach((el) =>
      this.listen(el, 'input', this.handleInput)
    )
  }
  disconnectedCallback() {
    this.unlistenAll()
  }
}
Object.assign(EditLinkForm.prototype, renderDataMixin)
customElements.define('edit-link-form', EditLinkForm, { extends: 'form' })
