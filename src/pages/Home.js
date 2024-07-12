import React from 'react';
import { Link } from 'react-router-dom';
import CountUp from 'react-countup';
import Sidebar from '../components/Sidebar';
import image0 from '../assets/image0.png';
import IMG_3755 from '../assets/IMG_3755.png'
import IMG_5524 from '../assets/IMG_5524.png'
import unnamed from '../assets/unnamed.png' // Ensure the logo path is correct

import '../styles/home.css'; // Ensure you have this CSS file

function Home() {
    return (
        <body className="font-kanit text-lg text-white bg-gray-800 dark:bg-gray-900">
            <main>
                <section className="relative pt-20 pb-8 personal-wrapper overflow-hidden bg-purple-500/5 text-center" id="home">
                    <div className="container mx-auto">
                        <div className="flex flex-col items-center">
                            <img src={image0} alt="Bravus Logo" className="rounded-full shadow-md shadow-gray-200 dark:shadow-gray-800 mb-8" />
                            <h4 className="font-bold lg:text-[48px] text-4xl lg:leading-normal leading-normal mb-4 text-purple-500">
                                Welcome to
                                <span className="index-module_type__E-SaG typewrite"> Bravus</span>
                            </h4>
                            <p className="text-gray-400 max-w-xl mx-auto lg:text-[20px]">
                                Your premier seamless appointment scheduling.
                            </p>
                            <div className="mt-6">
                                <button className="btn bg-purple-500 hover:bg-purple-600 border-purple-500 hover:border-purple-600 text-white rounded-md" onClick={() => document.getElementById('about').scrollIntoView({ behavior: 'smooth' })}>Learn More</button>
                            </div>
                        </div>
                    </div>
                </section>

                <section id="about" className="relative md:py-24 py-16 bg-gray-700 text-center">
                    <div className="container mx-auto">
                        <div className="grid md:grid-cols-1 grid-cols-1 items-center gap-[30px]">
                            <div>
                                <h2 className="lg:text-[48px] font-bold mb-4 text-purple-500">Why Bravus?</h2>
                                <p className="text-gray-400 max-w-xl mx-auto lg:text-[20px]">
                                    We aimed to create an app specifically for service-based businesses, such as salons and barbershops, to streamline their operations seamlessly. Our app integrates with the Google Sheets API to manage data efficiently and features an inbuilt appointment system that allows users to schedule and track appointments effortlessly. The app also includes advanced graphics and tools to help business owners manage their schedules and track their performance intuitively.
                                </p>
                            </div>
                        </div>
                    </div>
                </section>

                <section id="Team" className="relative md:py-24 py-16 bg-gray-800 text-center">
                    <div className="container mx-auto">
                        <div className="grid grid-cols-1">
                            <h2 className="lg:text-[48px] font-bold mb-4 text-purple-500">The Team</h2>
                            <p className="text-gray-400 max-w-xl mx-auto mb-8 lg:text-[20px]">
                                Meet the talented team behind Bravus. Together, we have brought our final project for our Software Engineering certificate to life, combining our skills and expertise to create an exceptional appointment scheduling experience.
                            </p>
                            <div className="grid md:grid-cols-3 gap-8">
                                <div className="team-member">
                                    <div className="team-member-photo">
                                        <img src={IMG_5524} alt="Sean A. Cardona" />
                                        <div className="peel-effect"></div>
                                    </div>
                                    <p className="team-member-name">Sean A. Cardona</p>
                                </div>
                                <div className="team-member">
                                    <div className="team-member-photo">
                                        <img src={unnamed} alt="Aramis Martinez" />
                                        <div className="peel-effect"></div>
                                    </div>
                                    <p className="team-member-name">Aramis Martinez</p>
                                </div>
                                <div className="team-member">
                                    <div className="team-member-photo">
                                        <img src={IMG_3755} alt="Yeneishla Santiago" />
                                        <div className="peel-effect"></div>
                                    </div>
                                    <p className="team-member-name">Yeneishla Santiago</p>
                                </div>
                            </div>
                        </div>
                    </div>
                </section>

            </main>
        </body>
    );
}

export default Home;