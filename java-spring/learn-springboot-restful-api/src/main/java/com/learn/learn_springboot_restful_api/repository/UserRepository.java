package com.learn.learn_springboot_restful_api.repository;

import com.learn.learn_springboot_restful_api.entity.User;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface UserRepository extends JpaRepository<User, String> {
}
