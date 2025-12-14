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
          <form class="contact-form" @submit.prevent="submitForm" name="contact">
            <!-- Hidden fields for Netlify -->
            <input type="hidden" name="form-name" value="contact" />
            <input type="hidden" name="bot-field" v-model="form.botField" />
            
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
                  name="phone"
                  type="tel" 
                  placeholder="555-123-4567"
                  @input="formatPhoneInput"
                  maxlength="12"
                  :disabled="isSubmitting" 
                />
              </div>
              <div class="field-group">
                <label for="location">Project Location (City / County)</label>
                <input 
                  id="location" 
                  v-model="form.location" 
                  name="location"
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
                  name="timeline"
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
              <div class="budget-input-wrapper">
                <span class="budget-prefix">$</span>
                <input 
                  id="budget" 
                  v-model="form.budget" 
                  name="budget"
                  type="text" 
                  placeholder="250,000" 
                  @input="formatBudgetInput"
                  :disabled="isSubmitting" 
                />
              </div>
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
                name="files"
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
        botField: '', // honeypot field
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
    
    formatPhoneInput(event) {
      // Remove all non-numeric characters
      let value = event.target.value.replace(/\D/g, '')
      
      // Limit to 10 digits
      value = value.substring(0, 10)
      
      // Format as XXX-XXX-XXXX
      if (value.length >= 6) {
        value = `${value.slice(0, 3)}-${value.slice(3, 6)}-${value.slice(6)}`
      } else if (value.length >= 3) {
        value = `${value.slice(0, 3)}-${value.slice(3)}`
      }
      
      this.form.phone = value
    },
    
    formatBudgetInput(event) {
      // Remove all non-numeric characters except commas
      let value = event.target.value.replace(/[^\d,]/g, '')
      
      // Remove existing commas
      value = value.replace(/,/g, '')
      
      // Add commas for thousands
      if (value) {
        value = parseInt(value, 10).toLocaleString('en-US')
      }
      
      this.form.budget = value
    },
    
    async submitForm() {
      this.isSubmitting = true
      this.formStatus = { message: '', type: '' }
      
      try {
        // Encode form data for Netlify
        const encode = (data) => {
          return Object.keys(data)
            .map(key => encodeURIComponent(key) + "=" + encodeURIComponent(data[key]))
            .join("&")
        }
        
        // Prepare form data
        const formData = {
          'form-name': 'contact',
          'bot-field': this.form.botField,
          'name': this.form.name,
          'email': this.form.email,
          'phone': this.form.phone,
          'location': this.form.location,
          'projectType': this.form.projectType,
          'timeline': this.form.timeline,
          'budget': this.form.budget,
          'message': this.form.message
        }
        
        // Submit to Netlify Forms
        const response = await fetch('/', {
          method: 'POST',
          headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
          body: encode(formData)
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
        botField: '',
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

.budget-input-wrapper {
  position: relative;
  display: flex;
  align-items: center;
}

.budget-prefix {
  position: absolute;
  left: 12px;
  color: #333;
  font-weight: 500;
  pointer-events: none;
  z-index: 1;
}

.budget-input-wrapper input {
  padding-left: 28px;
  width: 100%;
}
</style>