import { BaseForm } from './base-form'
import { camelize, kebabize } from '../../lib/utils'

function castType(value, type) {
  switch (type) {
    case 'int':
      return parseInt(value)
    default:
      return value
  }
}

export class StateManager extends BaseForm {
  constructor() {
    super()
    this.setDisabledButtons()
  }

  set disabled(bool) {
    if (bool) {
      Object.values(this.inputs).forEach((el) => (el.disabled = true))
      Object.values(this.buttons).forEach((el) => (el.disabled = true))
    } else {
      Object.values(this.inputs).forEach((el) => (el.disabled = false))
      Object.values(this.buttons).forEach((el) => (el.disabled = false))
      this.setDisabledButtons()
    }
  }

  get resourceId() {
    return parseInt(this.inputs['id'].value)
  }

  set resourceId(v) {
    this.inputs['id'].value = v
  }

  set state(d) {
    // we are storing set data here to compare with state
    this._data = d
    let el
    for (const prop in d) {
      el = this.inputs[kebabize(prop)]
      if (el) {
        switch (el.type) {
          case 'checkbox':
            el.checked = !!d[prop]
            break
          default:
            el.value = d[prop]
            break
        }
      }
    }
  }
  get state() {
    const data = {}
    let el
    for (const name in this.inputs) {
      el = this.inputs[name]
      switch (el.type) {
        case 'checkbox':
          data[camelize(name)] = el.checked
          break
        default:
          data[camelize(name)] = castType(
            el.value,
            el.getAttribute('data-type')
          )
          break
      }
    }
    return data
  }

  get hasChanges() {
    if (!this._data) return false
    const state = this.state
    for (const prop in this._data) {
      if (state.hasOwnProperty(prop) && this._data[prop] !== state[prop]) {
        return true
      }
    }
    return false
  }

  setDisabledButtons() {
    this.querySelectorAll('button[data-require-changes]').forEach((el) => {
      el.disabled = !this.hasChanges
    })
  }

  connectedCallback() {
    super.connectedCallback()

    Object.values(this.inputs).forEach((el) =>
      this.listen(el, 'input', this.setDisabledButtons)
    )
  }
  disconnectedCallback() {
    this.unlistenAll()
  }
}
customElements.define('state-manager', StateManager, { extends: 'form' })
