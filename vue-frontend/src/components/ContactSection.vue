<template>
  <section id="contact">
    <div class="container">
      <div class="section-heading">
        <div class="section-eyebrow">Contact & Quote</div>
        <h2>Tell FALTOR Consulting about your project.</h2>
        <p>
          Use the form below to request a quote, share your project details, and upload any sketches or inspiration.
          You'll receive a follow-up to discuss scope, pricing, and timeline.
        </p>
      </div>

      <div id="quote" class="contact-grid">
        <div class="card">
          <form class="contact-form" @submit.prevent="submitForm" name="contact" method="POST" netlify netlify-honeypot="bot-field">
            <!-- Hidden fields for Netlify -->
            <input type="hidden" name="form-name" value="contact" />
            <input type="hidden" name="bot-field" />
            
            <div v-if="formStatus.message" 
                 :class="['form-message', formStatus.type]">
              {{ formStatus.message }}
            </div>
            
            <div class="inline-fields">
              <div class="field-group">
                <label for="name">Name *</label>
                <input 
                  id="name" 
                  v-model="form.name" 
                  name="name"
                  type="text" 
                  required 
                  :disabled="isSubmitting" 
                />
              </div>
              <div class="field-group">
                <label for="email">Email *</label>
                <input 
                  id="email" 
                  v-model="form.email" 
                  name="email"
                  type="email" 
                  required 
                  :disabled="isSubmitting" 
                />
              </div>
            </div>

            <div class="inline-fields">
              <div class="field-group">
                <label for="phone">Phone</label>
                <input 
                  id="phone" 
                  v-model="form.phone" 
                  type="tel" 
                  :disabled="isSubmitting" 
                />
              </div>
              <div class="field-group">
                <label for="location">Project Location (City / County)</label>
                <input 
                  id="location" 
                  v-model="form.location" 
                  type="text" 
                  placeholder="e.g., Goldsboro, Eastern NC" 
                  :disabled="isSubmitting" 
                />
              </div>
            </div>

            <div class="inline-fields">
              <div class="field-group">
                <label for="projectType">Project Type *</label>
                <select 
                  id="projectType" 
                  v-model="form.projectType" 
                  name="projectType"
                  required 
                  :disabled="isSubmitting"
                >
                  <option value="">Select one</option>
                  <option>Custom house plan</option>
                  <option>Interior remodel</option>
                  <option>Addition / porch</option>
                  <option>Garage / accessory building</option>
                  <option>As-built drawings</option>
                  <option>Site / plot plan</option>
                  <option>3D rendering only</option>
                  <option>Other drafting services</option>
                </select>
              </div>
              <div class="field-group">
                <label for="timeline">Desired Start Timeline</label>
                <select 
                  id="timeline" 
                  v-model="form.timeline" 
                  :disabled="isSubmitting"
                >
                  <option value="">Select one</option>
                  <option>As soon as possible</option>
                  <option>Within 1–2 months</option>
                  <option>Within 3–6 months</option>
                  <option>Flexible / just planning</option>
                </select>
              </div>
            </div>

            <div class="field-group">
              <label for="budget">Approximate Construction Budget (Optional)</label>
              <input 
                id="budget" 
                v-model="form.budget" 
                type="text" 
                placeholder="e.g., $250,000 – $400,000" 
                :disabled="isSubmitting" 
              />
              <p class="hint">This simply helps align the design with realistic construction costs.</p>
            </div>

            <div class="field-group">
              <label for="message">Project Details *</label>
              <textarea 
                id="message" 
                v-model="form.message" 
                name="message"
                required 
                placeholder="Describe your project, goals, and any important details."
                :disabled="isSubmitting"
              ></textarea>
            </div>

            <div class="field-group upload-field">
              <label for="sketches">Upload Your Sketches / Files</label>
              <input 
                id="sketches" 
                type="file" 
                multiple 
                @change="handleFileUpload"
                :disabled="isSubmitting" 
              />
              <p class="hint">
                Attach scans, photos of hand sketches, or any existing plans.
              </p>
            </div>

            <div class="field-group">
              <button 
                type="submit" 
                class="btn btn-primary"
                :disabled="isSubmitting"
              >
                {{ isSubmitting ? 'Submitting...' : 'Submit Quote Request' }}
              </button>
              <p class="hint">
                Your request will be sent directly to FALTOR Consulting.
              </p>
            </div>
          </form>
        </div>

        <aside class="contact-side">
          <div class="contact-block">
            <strong>Contact Information</strong>
            <ul class="contact-list">
              <li>Phone: <a href="tel:2525217965">252-521-7965</a></li>
              <li>Email: <a href="mailto:faltorco.ben@gmail.com">faltorco.ben@gmail.com</a></li>
            </ul>
            <p class="hint">
              Serving homeowners, builders, and investors throughout Eastern North Carolina.
            </p>
          </div>

          <div class="contact-block">
            <strong>Typical Hours</strong>
            <p>Monday – Friday: 8:00 AM – 5:00 PM</p>
            <p class="hint">Evening consultations available by appointment.</p>
          </div>
        </aside>
      </div>
    </div>
  </section>
</template>

<script>
import axios from 'axios'

export default {
  name: 'ContactSection',
  data() {
    return {
      isSubmitting: false,
      formStatus: {
        message: '',
        type: ''
      },
      form: {
        name: '',
        email: '',
        phone: '',
        location: '',
        projectType: '',
        timeline: '',
        budget: '',
        message: ''
      },
      files: []
    }
  },
  methods: {
    handleFileUpload(event) {
      this.files = Array.from(event.target.files)
    },
    
    async submitForm() {
      this.isSubmitting = true
      this.formStatus = { message: '', type: '' }
      
      try {
        const formData = new FormData()
        
        // Add form fields
        Object.keys(this.form).forEach(key => {
          formData.append(key, this.form[key])
        })
        
        // Add files
        this.files.forEach(file => {
          formData.append('files', file)
        })
        
        // Add form name for Netlify
        formData.append('form-name', 'contact')
        
        // Submit to Netlify Forms
        const response = await fetch('/', {
          method: 'POST',
          body: formData
        })
        
        if (response.ok) {
          this.formStatus = {
            message: 'Thank you! Your message has been sent successfully. We will contact you soon!',
            type: 'success'
          }
          this.resetForm()
        } else {
          throw new Error('Form submission failed')
        }
      } catch (error) {
        console.error('Form submission error:', error)
        this.formStatus = {
          message: 'Sorry, there was an error sending your message. Please try again.',
          type: 'error'
        }
      } finally {
        this.isSubmitting = false
      }
    },
    
    resetForm() {
      this.form = {
        name: '',
        email: '',
        phone: '',
        location: '',
        projectType: '',
        timeline: '',
        budget: '',
        message: ''
      }
      this.files = []
      const fileInput = document.getElementById('sketches')
      if (fileInput) fileInput.value = ''
    }
  }
}
</script>

<style scoped>
.form-message {
  padding: 0.75rem 1rem;
  border-radius: 8px;
  margin-bottom: 1rem;
  font-size: 0.9rem;
}

.form-message.success {
  background-color: #d4edda;
  color: #155724;
  border: 1px solid #c3e6cb;
}

.form-message.error {
  background-color: #f8d7da;
  color: #721c24;
  border: 1px solid #f5c6cb;
}

.btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}
</style>