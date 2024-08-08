import React from 'react';
import { Link } from 'react-router-dom';
import Yene from '../assets/Yene.png';
import Sean from '../assets/Sean.png';
import Aramis from '../assets/Aramis.png'; // Ensure the logo path is correct

import '../styles/home.css'; // Ensure you have this CSS file

function Home() {
    return (
        <body className="font-kanit text-lg text-white bg-gray-800">
            <main>
                <section class="relative overflow-hidden md:py-48 py-40 bg-Purple/5" id="home">
                    <div class="container relative mt-8">
                        <div class="flex space-x-4 justify-between">
                            <div class="">
                                <h1 class="font-semibold lg:leading-normal leading-normal tracking-wide text-4xl lg:text-5xl mb-5">Get Creative &amp; Modern With Upcover</h1>
                                <p class="text-slate-400 text-lg max-w-xl">This is just a simple text made for this unique and awesome template, you can replace it with any text.</p>
                                <div class="mt-6">
                                    <a href="" class="h-10 px-6 tracking-wide inline-flex items-center justify-center font-medium rounded-md bg-teal-500 text-white">Contact Us <i class="mdi mdi-chevron-right ms-1"></i></a>
                                </div>
                            </div>
                            <div class="lg:ms-8">
                                <div class="relative">
                                    <img src="https://shreethemes.in/upcover/layouts/assets/images/design-team.svg" class="relative top-16" alt=""/>
                                    <div class="overflow-hidden absolute md:size-[500px] size-[400px] bg-gradient-to-tl to-teal-500/20 via-teal-500/70 from-teal-500 bottom-1/2 translate-y-1/2 md:start-0 start-1/2 ltr:md:translate-x-0 ltr:-translate-x-1/2 rtl:md:translate-x-0 rtl:translate-x-1/2 -z-1 shadow-md shadow-teal-500/10 rounded-full"></div>
                                </div>
                            </div>
                        </div>
                    </div>
                </section>sss

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

                <section id="Team" className="relative md:py-24 py-16 text-center">
                    <div className="container mx-auto">
                        <div className="grid grid-cols-1">
                            <h2 className="lg:text-[48px] font-bold mb-4 text-purple-500">The Team</h2>
                            <p className="text-gray-400 max-w-xl mx-auto mb-8 lg:text-[20px]">
                                Meet the talented team behind Bravus. Together, we have brought our final project for our Software Engineering certificate to life, combining our skills and expertise to create an exceptional appointment scheduling experience.
                            </p>
                            <div className="grid md:grid-cols-3 gap-8">
                                <div className="team-member">
                                    <div className="team-member-photo">
                                        <Link to="https://www.linkedin.com/in/sean-cardona-bb1594166?trk=contact-info">
                                            <img src={Sean} alt="Sean A. Cardona" />
                                        </Link>
                                        <div className="peel-effect"></div>
                                        <p className="team-member-name">Sean A. Cardona</p>
                                    </div>
                                </div>
                                <div className="team-member">
                                    <div className="team-member-photo">
                                        <Link to="https://www.linkedin.com/in/aramis-martinez-a1a507296/">
                                            <img src={Aramis} alt="Aramis Martinez" />
                                        </Link>
                                        <div className="peel-effect"></div>
                                        <p className="team-member-name">Aramis Martinez</p>
                                    </div>
                                </div>
                                <div className="team-member">
                                    <div className="team-member-photo">
                                        <Link to="https://www.linkedin.com/in/yeneishla-santiago-958b63254/">
                                            <img src={Yene} alt="Yeneishla Santiago" />
                                        </Link>
                                        <div className="peel-effect"></div>
                                        <p className="team-member-name">Yeneishla Santiago</p>
                                    </div>
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