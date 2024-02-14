import { ProgressiveBody } from './progressive-body'
class HomePage extends ProgressiveBody {}
customElements.define('home-page', HomePage, { extends: 'body' })
