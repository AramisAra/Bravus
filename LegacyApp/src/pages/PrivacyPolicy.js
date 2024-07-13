import React from 'react';

function PrivacyPolicy() {
  return (
    <div className="p-4 max-w-2xl mx-auto bg-gray-800 text-white rounded-lg shadow-lg">
      <h1 className="text-2xl font-bold mb-4">Privacy Policy</h1>
      <p className="mb-4">This Privacy Policy explains how we collect, use, disclose, and safeguard your information when you visit our website. Please read this privacy policy carefully. If you do not agree with the terms of this privacy policy, please do not access the site.</p>
      
      <h2 className="text-xl font-semibold mb-2">1. Information We Collect</h2>
      <p className="mb-4">We may collect information about you in a variety of ways. The information we may collect on the Site includes:</p>
      <ul className="list-disc list-inside mb-4">
        <li>Personal Data: Personally identifiable information, such as your name, shipping address, email address, and telephone number, and demographic information, such as your age, gender, hometown, and interests, that you voluntarily give to us when you register with the Site.</li>
        <li>Derivative Data: Information our servers automatically collect when you access the Site, such as your IP address, your browser type, your operating system, your access times, and the pages you have viewed directly before and after accessing the Site.</li>
        <li>Financial Data: Financial information, such as data related to your payment method (e.g., valid credit card number, card brand, expiration date) that we may collect when you purchase, order, return, exchange, or request information about our services from the Site.</li>
      </ul>

      <h2 className="text-xl font-semibold mb-2">2. Use of Your Information</h2>
      <p className="mb-4">Having accurate information about you permits us to provide you with a smooth, efficient, and customized experience. Specifically, we may use information collected about you via the Site to:</p>
      <ul className="list-disc list-inside mb-4">
        <li>Create and manage your account.</li>
        <li>Process your transactions and send you related information, including purchase confirmations and invoices.</li>
        <li>Administer sweepstakes, promotions, and contests.</li>
        <li>Request feedback and contact you about your use of the Site.</li>
        <li>Resolve disputes and troubleshoot problems.</li>
        <li>Respond to product and customer service requests.</li>
        <li>Deliver targeted advertising, newsletters, and other information regarding promotions and the Site to you.</li>
      </ul>

      <h2 className="text-xl font-semibold mb-2">3. Disclosure of Your Information</h2>
      <p className="mb-4">We may share information we have collected about you in certain situations. Your information may be disclosed as follows:</p>
      <ul className="list-disc list-inside mb-4">
        <li>By Law or to Protect Rights: If we believe the release of information about you is necessary to respond to legal process, to investigate or remedy potential violations of our policies, or to protect the rights, property, and safety of others.</li>
        <li>Business Transfers: We may share or transfer your information in connection with, or during negotiations of, any merger, sale of company assets, financing, or acquisition of all or a portion of our business to another company.</li>
        <li>Affiliates: We may share your information with our affiliates, in which case we will require those affiliates to honor this Privacy Policy.</li>
        <li>Business Partners: We may share your information with our business partners to offer you certain products, services, or promotions.</li>
      </ul>

      <h2 className="text-xl font-semibold mb-2">4. Security of Your Information</h2>
      <p className="mb-4">We use administrative, technical, and physical security measures to help protect your personal information. While we have taken reasonable steps to secure the personal information you provide to us, please be aware that despite our efforts, no security measures are perfect or impenetrable, and no method of data transmission can be guaranteed against any interception or other type of misuse.</p>

      <h2 className="text-xl font-semibold mb-2">5. Contact Us</h2>
      <p>If you have questions or comments about this Privacy Policy, please contact us at:</p>
      <p className="mb-4">
        <strong>Email:</strong> support@bravus.com
      </p>
    </div>
  );
}

export default PrivacyPolicy;
