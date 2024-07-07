import React from 'react';

function Licensing() {
  return (
    <div className="p-4 max-w-2xl mx-auto bg-gray-800 text-white rounded-lg shadow-lg">
      <h1 className="text-2xl font-bold mb-4">Licensing Information</h1>
      <p className="mb-4">This Licensing Information explains the terms and conditions under which our software and services are licensed. Please read this licensing information carefully. By using our software and services, you agree to be bound by these terms.</p>

      <h2 className="text-xl font-semibold mb-2">1. License Grant</h2>
      <p className="mb-4">We grant you a non-exclusive, non-transferable, revocable license to access and use our software and services strictly in accordance with these terms of use.</p>

      <h2 className="text-xl font-semibold mb-2">2. Restrictions</h2>
      <p className="mb-4">You agree not to, and you will not permit others to:</p>
      <ul className="list-disc list-inside mb-4">
        <li>License, sell, rent, lease, assign, distribute, transmit, host, outsource, disclose, or otherwise commercially exploit the service or make the platform available to any third party.</li>
        <li>Modify, make derivative works of, disassemble, decrypt, reverse compile, or reverse engineer any part of the service.</li>
        <li>Remove, alter, or obscure any proprietary notice (including any notice of copyright or trademark) of ours or its affiliates, partners, suppliers, or the licensors of the service.</li>
      </ul>

      <h2 className="text-xl font-semibold mb-2">3. Modifications to Service</h2>
      <p className="mb-4">We reserve the right to modify, suspend, or discontinue, temporarily or permanently, the service or any part thereof with or without notice.</p>

      <h2 className="text-xl font-semibold mb-2">4. Termination</h2>
      <p className="mb-4">We may terminate or suspend your account and bar access to the service immediately, without prior notice or liability, under our sole discretion, for any reason whatsoever and without limitation, including but not limited to a breach of the terms.</p>

      <h2 className="text-xl font-semibold mb-2">5. Governing Law</h2>
      <p className="mb-4">These terms shall be governed and construed in accordance with the laws of [Your Country], without regard to its conflict of law provisions.</p>

      <h2 className="text-xl font-semibold mb-2">6. Contact Us</h2>
      <p>If you have any questions about these terms, please contact us at:</p>
      <p className="mb-4">
        <strong>Email:</strong> support@bravus.com
      </p>
    </div>
  );
}

export default Licensing;
