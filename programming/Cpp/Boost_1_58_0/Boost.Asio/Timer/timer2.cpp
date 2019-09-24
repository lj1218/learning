// Timer.2 - Using a timer asynchronously
// https://www.boost.org/doc/libs/1_58_0/doc/html/boost_asio/tutorial/tuttimer2.html
#include <iostream>
#include <boost/asio.hpp>
#include <boost/date_time/posix_time/posix_time.hpp>

//void print(const boost::system::error_code& /*e*/)
void print(const boost::system::error_code& e)
{
  std::cout << "Hello, world! " << e << std::endl;
}

int main()
{
  boost::asio::io_service io;

  boost::asio::deadline_timer t(io, boost::posix_time::seconds(5));
  t.async_wait(&print);

  io.run();

  return 0;
}

// g++ -std=c++11 -lboost_system timer2.cpp -o timer2
