
#ifndef GOVALHALLA_ACTOR_H
#define GOVALHALLA_ACTOR_H

#include <memory>
#include <stdexcept>
#include <sstream>
#include <boost/property_tree/ptree.hpp>
#include <boost/property_tree/json_parser.hpp>
#include <valhalla/tyr/actor.h>
#include "valhalla/worker.h"  // For valhalla_exception_t
#include <valhalla/baldr/graphreader.h>
#include <valhalla/proto/api.pb.h>



namespace govalhalla {

class Actor {
public:

    explicit  Actor(const std::string& config_json, bool auto_cleanup = true);

    std::pair<std::string, std::string> Act(valhalla::Api& api, const std::function<void()>* interrupt = nullptr);
    std::pair<std::string, std::string> Route(const std::string& request, const std::function<void()>* interrupt = nullptr, valhalla::Api* api = nullptr);
    std::pair<std::string, std::string> Locate(const std::string& request, const std::function<void()>* interrupt = nullptr, valhalla::Api* api = nullptr);
    std::pair<std::string, std::string> Matrix(const std::string& request, const std::function<void()>* interrupt = nullptr, valhalla::Api* api = nullptr);
    std::pair<std::string, std::string> OptimizedRoute(const std::string& request, const std::function<void()>* interrupt = nullptr, valhalla::Api* api = nullptr);
    std::pair<std::string, std::string> Isochrone(const std::string& request, const std::function<void()>* interrupt = nullptr, valhalla::Api* api = nullptr);
    std::pair<std::string, std::string> TraceRoute(const std::string& request, const std::function<void()>* interrupt = nullptr, valhalla::Api* api = nullptr);
    std::pair<std::string, std::string> TraceAttributes(const std::string& request, const std::function<void()>* interrupt = nullptr, valhalla::Api* api = nullptr);
    std::pair<std::string, std::string> Height(const std::string& request, const std::function<void()>* interrupt = nullptr, valhalla::Api* api = nullptr);
    std::pair<std::string, std::string> TransitAvailable(const std::string& request, const std::function<void()>* interrupt = nullptr, valhalla::Api* api = nullptr);
    std::pair<std::string, std::string> Expansion(const std::string& request, const std::function<void()>* interrupt = nullptr, valhalla::Api* api = nullptr);
    std::pair<std::string, std::string> Centroid(const std::string& request, const std::function<void()>* interrupt = nullptr, valhalla::Api* api = nullptr);
    std::pair<std::string, std::string> Status(const std::string& request, const std::function<void()>* interrupt = nullptr, valhalla::Api* api = nullptr);

    void Cleanup();
    // Destructor
    ~Actor() = default;

private:
    boost::property_tree::ptree parseConfig(const std::string& config_json);
    std::string exception_to_json(const valhalla::valhalla_exception_t &ex);

    std::unique_ptr<valhalla::tyr::actor_t> actor_;
};

} // namespace govalhalla

#endif // GOVALHALLA_ACTOR_H