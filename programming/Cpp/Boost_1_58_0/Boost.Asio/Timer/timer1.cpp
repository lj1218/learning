// Timer.1 - Using a timer synchronously
// https://www.boost.org/doc/libs/1_58_0/doc/html/boost_asio/tutorial/tuttimer1.html
#include <iostream>
#include <boost/asio.hpp>
#include <boost/date_time/posix_time/posix_time.hpp>

int main()
{
    boost::asio::io_service io;
    boost::asio::deadline_timer t(io, boost::posix_time::seconds(5));
    t.wait();
    std::cout << "Hello, world!" << std::endl;
    return 0;
}

// g++ -std=c++11 -lboost_system timer1.cpp -o timer1
