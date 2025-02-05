
#include <memory>
#include <sstream>
#include <functional>
#include <stdexcept>
#include "boost/property_tree/json_parser.hpp"
#include "valhalla/tyr/actor.h"
#include "valhalla/baldr/graphreader.h"
#include "valhalla/proto/api.pb.h"
#include "boost/property_tree/ptree.hpp"
#include "valhalla/baldr/rapidjson_utils.h"
#include "govalhalla_actor.h"



namespace govalhalla {

Actor::Actor(const std::string& config_json, bool auto_cleanup) 
        : actor_(std::make_unique<valhalla::tyr::actor_t>(parseConfig(config_json), auto_cleanup)) {}



std::pair<std::string, std::string> Actor::Act(valhalla::Api& api, const std::function<void()>* interrupt) {
    try {
        return std::make_pair(actor_->act(api, interrupt), "");
        } catch (const std::exception& e) {
            return std::make_pair("", std::string("Error: ") + e.what());
        } catch (...) {
            return std::make_pair("", "Unknown error");
        }
        
        
    }

std::pair<std::string, std::string> Actor::Route(const std::string& request, const std::function<void()>* interrupt, valhalla::Api* api) {
    try {
        return std::make_pair(actor_->route(request, interrupt, api), "");
            
        } catch (const std::exception& e) {
            return std::make_pair("", std::string("Error: ") + e.what());
        } catch (...) {
            return std::make_pair("", "Unknown error");
        }
        
        
    }

std::pair<std::string, std::string> Actor::Locate(const std::string& request, const std::function<void()>* interrupt, valhalla::Api* api) {
    try {
        return std::make_pair(actor_->locate(request, interrupt, api), "");
            
        } catch (const std::exception& e) {
            return std::make_pair("", std::string("Error: ") + e.what());
        } catch (...) {
            return std::make_pair("", "Unknown error");
        }
        
        
    }

std::pair<std::string, std::string> Actor::Matrix(const std::string& request, const std::function<void()>* interrupt, valhalla::Api* api) {
    try {
        return std::make_pair(actor_->matrix(request, interrupt, api), "");
            
            
        } catch (const std::exception& e) {
            return std::make_pair("", std::string("Error: ") + e.what());
        } catch (...) {
            return std::make_pair("", "Unknown error");
        }
        
        
    }

std::pair<std::string, std::string> Actor::OptimizedRoute(const std::string& request, const std::function<void()>* interrupt, valhalla::Api* api) {
    try {
        return std::make_pair(actor_->optimized_route(request, interrupt, api), "");
            
            
        } catch (const std::exception& e) {
            return std::make_pair("", std::string("Error: ") + e.what());
        } catch (...) {
            return std::make_pair("", "Unknown error");
        }
        
        
    }

std::pair<std::string, std::string> Actor::Isochrone(const std::string& request, const std::function<void()>* interrupt, valhalla::Api* api) {
    try {
        return std::make_pair(actor_->isochrone(request, interrupt, api), "");
            
            
        } catch (const std::exception& e) {
            return std::make_pair("", std::string("Error: ") + e.what());
        } catch (...) {
            return std::make_pair("", "Unknown error");
        }
        
        
    }


std::pair<std::string, std::string> Actor::TraceRoute(const std::string& request, const std::function<void()>* interrupt, valhalla::Api* api) {
    try {
        return std::make_pair(actor_->trace_route(request, interrupt, api), "");
            
            
        } catch (const std::exception& e) {
            return std::make_pair("", std::string("Error: ") + e.what());
        } catch (...) {
            return std::make_pair("", "Unknown error");
        }
        
        
    }


std::pair<std::string, std::string> Actor::TraceAttributes(const std::string& request, const std::function<void()>* interrupt, valhalla::Api* api) {
    try {
        return std::make_pair(actor_->trace_attributes(request, interrupt, api), "");
            
            
        } catch (const std::exception& e) {
            return std::make_pair("", std::string("Error: ") + e.what());
        } catch (...) {
            return std::make_pair("", "Unknown error");
        }
        
        
    }


std::pair<std::string, std::string> Actor::Height(const std::string& request, const std::function<void()>* interrupt, valhalla::Api* api) {
    try {
        return std::make_pair(actor_->height(request, interrupt, api), "");
            
            
        } catch (const std::exception& e) {
            return std::make_pair("", std::string("Error: ") + e.what());
        } catch (...) {
            return std::make_pair("", "Unknown error");
        }
        
        
    }


std::pair<std::string, std::string> Actor::TransitAvailable(const std::string& request, const std::function<void()>* interrupt, valhalla::Api* api) {
    try {
        return std::make_pair(actor_->transit_available(request, interrupt, api), "");
            
            
        } catch (const std::exception& e) {
            return std::make_pair("", std::string("Error: ") + e.what());
        } catch (...) {
            return std::make_pair("", "Unknown error");
        }
        
        
    }


std::pair<std::string, std::string> Actor::Expansion(const std::string& request, const std::function<void()>* interrupt, valhalla::Api* api) {
    try {
        return std::make_pair(actor_->expansion(request, interrupt, api), "");
            
            
        } catch (const std::exception& e) {
            return std::make_pair("", std::string("Error: ") + e.what());
        } catch (...) {
            return std::make_pair("", "Unknown error");
        }
        
        
    }


std::pair<std::string, std::string> Actor::Centroid(const std::string& request, const std::function<void()>* interrupt, valhalla::Api* api) {
    try {
        return std::make_pair(actor_->centroid(request, interrupt, api), "");
            
            
        } catch (const std::exception& e) {
            return std::make_pair("", std::string("Error: ") + e.what());
        } catch (...) {
            return std::make_pair("", "Unknown error");
        }
        
        
    }


std::pair<std::string, std::string> Actor::Status(const std::string& request, const std::function<void()>* interrupt, valhalla::Api* api) {
    try {
        return std::make_pair(actor_->status(request, interrupt, api), "");
            
            
        } catch (const std::exception& e) {
            return std::make_pair("", std::string("Error: ") + e.what());
        } catch (...) {
            return std::make_pair("", "Unknown error");
        }
        
        
    }

    
void Actor::Cleanup() {
    if (actor_) {
            actor_->cleanup();
        }
}

boost::property_tree::ptree Actor::parseConfig(const std::string& config_json) {

        boost::property_tree::ptree pt;
        std::istringstream is(config_json);
        rapidjson::read_json(is, pt);
        return pt;
    }

} // namespace govalhalla
