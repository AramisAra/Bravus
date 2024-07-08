import React from 'react';

const LandingPage = () => {
  return (
    <div className="min-h-screen bg-gray-900 text-white flex flex-col">
      <header className="bg-gray-800 p-4 flex justify-between items-center">
        <h1 className="text-3xl font-bold">Bravus<span className="text-blue-500">.</span></h1>
        <nav className="flex space-x-4">
          <a href="#who-we-are" className="text-white hover:text-gray-400">Who We Are</a>
          <a href="#why-choose-us" className="text-white hover:text-gray-400">Why Choose Us</a>
        </nav>
      </header>

      <main className="flex-grow p-8">
        <section id="who-we-are" className="text-center mb-16">
          <h2 className="text-2xl font-bold mb-4">Who We Are</h2>
          <p className="mb-8">
            Welcome to Bravus, your premier destination for seamless appointment scheduling.
          </p>
          <img src="your-logo-url" alt="Bravus Logo" className="mx-auto mb-8" />
        </section>

        <section id="why-choose-us" className="text-center">
          <h2 className="text-2xl font-bold mb-4">Why We Chose This Project</h2>
          <p>
            This project is our final project for our Software Engineering certificate. We are a team of three consisting of Aramis, Sean, and Yeneishla.
          </p>
        </section>
      </main>

      <footer className="bg-gray-800 p-4 text-center">
        <p>&copy; 2024 Bravus. All rights reserved.</p>
      </footer>
    </div>
  );
};

export default LandingPage;
